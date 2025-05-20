package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DataBase interface {
	CheckExistsUser(login string) (bool, error)
	GetUserBalans(login string) (int64, error)
	GetPassword(login string) (string, error)
	CreateUser(login, password string) error
	Deposite(login string, cash int64) error
	TransferMoney(loginFrom, loginTo string, cash int64) error

	Close()
}

type pgxdb struct {
	pool *pgxpool.Pool
}

func NewDBConn(conn string, schemaName string, tableName string) (DataBase, error) {
	pool, err := pgxpool.New(context.Background(), conn)
	if err != nil {
		return nil, fmt.Errorf("can't create pool: %w", err)
	}
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("db didn't pinged: %w", err)
	}

	if err = createSchema(pool); err != nil {
		return nil, err
	}

	if err = createTables(pool); err != nil {
		return nil, err
	}

	if err = createTriggerFunc(pool); err != nil {
		return nil, err
	}

	if err = createTrigger(pool); err != nil {
		return nil, err
	}

	return &pgxdb{pool}, nil
}

func createSchema(pool *pgxpool.Pool) error {
	query := fmt.Sprintf("create schema if not exists %s", catalog[1].schema)
	if _, err := pool.Exec(context.Background(), query); err != nil {
		return fmt.Errorf("can't create schema: %w", err)
	}
	return nil
}

func createTables(pool *pgxpool.Pool) error {
	query := fmt.Sprintf(
		"create table if not exists %s.%s (\n"+
			"user_id bigserial primary key,\n"+
			"login   text not null unique,\n"+
			"password text not null\n"+
			");", catalog[1].schema, catalog[1].table)
	if _, err := pool.Exec(context.Background(), query); err != nil {
		return fmt.Errorf("can't create table %s: %w", catalog[1].table, err)
	}

	query = fmt.Sprintf(
		"create table if not exists %s.%s (\n"+
			"user_id bigint primary key\n"+
			"references %s.%s(user_id) on delete cascade,\n"+
			"cash bigint\n"+
			");", catalog[1].schema, catalog[2].table, catalog[1].schema, catalog[1].table)
	if _, err := pool.Exec(context.Background(), query); err != nil {
		return fmt.Errorf("can't create table %s: %w", catalog[2].table, err)
	}
	return nil
}

func createTriggerFunc(pool *pgxpool.Pool) error {
	query := fmt.Sprintf(
		"create or replace function %s.create_wallet()\n"+
			"returns trigger as $$\n"+
			"begin\n"+
			"insert into %s.%s(user_id, cash) values(NEW.user_id, 0);\n"+
			"return new;\n"+
			"end;"+
			"$$ language plpgsql;", catalog[1].schema, catalog[1].schema, catalog[2].table,
	)
	if _, err := pool.Exec(context.Background(), query); err != nil {
		return fmt.Errorf("can't trigger func %s: %w", catalog[2].table, err)
	}
	return nil
}

func createTrigger(pool *pgxpool.Pool) error {
	query := fmt.Sprintf(
		"create or replace trigger after_insert_user\n"+
			"after insert on %s.%s\n"+
			"for each row\n"+
			"execute function %s.create_wallet()", catalog[1].schema, catalog[1].table, catalog[1].schema,
	)
	if _, err := pool.Exec(context.Background(), query); err != nil {
		return fmt.Errorf("can't trigger %s: %w", catalog[2].table, err)
	}
	return nil
}

// METHOD OF INTERFACES

func (db *pgxdb) Close() {
	db.pool.Close()
}

func (db *pgxdb) CheckExistsUser(login string) (bool, error) {
	existQuery := fmt.Sprintf("select exists(select login from %s.%s where login = $1)", catalog[1].schema, catalog[1].table)

	var exists bool
	row := db.pool.QueryRow(context.Background(), existQuery, login)
	err := row.Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("problem with check: %w", err)
	}
	return exists, err
}

func (db *pgxdb) GetPassword(login string) (string, error) {
	query := fmt.Sprintf("select password from %s.%s where login = $1", catalog[1].schema, catalog[1].table)

	var hashPass string
	row := db.pool.QueryRow(context.Background(), query, login)
	err := row.Scan(&hashPass)
	if err != nil {
		return "", fmt.Errorf("can't get password: %w", err)
	}
	return hashPass, nil
}
func (db *pgxdb) CreateUser(login, password string) error {
	createUserQuery := fmt.Sprintf("insert into %s.%s (login, password) values ($1, $2)", catalog[1].schema, catalog[1].table)
	if _, err := db.pool.Exec(context.Background(), createUserQuery, login, password); err != nil {
		return err
	}

	return nil
}

func (db *pgxdb) Deposite(login string, cash int64) error {
	query := fmt.Sprintf(
		"update %s.%s "+
			"set cash = cash + $1 "+
			"where user_id = ("+
			"select user_id from %s.%s "+
			"where login = $2"+
			")", catalog[2].schema, catalog[2].table, catalog[1].schema, catalog[1].table,
	)
	if _, err := db.pool.Exec(context.Background(), query, cash, login); err != nil {
		return err
	}
	return nil
}

func (db *pgxdb) GetUserBalans(login string) (int64, error) {
	query := fmt.Sprintf(
		"select cash from %s.%s "+
			"where user_id = ("+
			"select user_id from %s.%s "+
			"where login = $1"+
			")", catalog[2].schema, catalog[2].table, catalog[1].schema, catalog[1].table,
	)
	var cash int64
	row := db.pool.QueryRow(context.Background(), query, login)
	if err := row.Scan(&cash); err != nil {
		return 0, fmt.Errorf("can't get a balance: %w", err)
	}
	return cash, nil
}

func (db *pgxdb) TransferMoney(loginFrom, loginTo string, cash int64) error {

	ts, err := db.pool.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("can't create transaction: %w", err)
	}

	defer func() {
		if err != nil {
			err = ts.Rollback(context.Background())
			if err != nil {
				err = fmt.Errorf("fatal error with Rollback: %w", err)
			}
		}
	}()

	err = db.Deposite(loginFrom, -cash)
	if err != nil {
		return fmt.Errorf("problem with deposit from %s: %w", loginFrom, err)
	}

	err = db.Deposite(loginTo, cash)
	if err != nil {
		return fmt.Errorf("problem with deposit to %s: %w", loginFrom, err)
	}

	// Auto commit if doen't have errors

	return nil
}
