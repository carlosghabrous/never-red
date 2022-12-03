CREATE TABLE IF NOT EXISTS mov_type (
    id serial PRIMARY KEY,
    name VARCHAR(10)
);

CREATE TABLE IF NOT EXISTS category (
    id serial PRIMARY KEY,
    name VARCHAR(50)
);
CREATE TABLE IF NOT EXISTS currency (
    id serial PRIMARY KEY,
    name VARCHAR(3)
);

CREATE TABLE IF NOT EXISTS movements (
    id SERIAL PRIMARY KEY,
    mov_date DATE NOT NULL,
    mov_type SMALLINT, 
    commerce VARCHAR(50),
    category SMALLINT,
    currency SMALLINT,
    amount REAL, 

    CONSTRAINT fk_movtype FOREIGN KEY(mov_type) REFERENCES mov_type(id) ON DELETE SET NULL,
    CONSTRAINT fk_category FOREIGN KEY (category) REFERENCES category(id) ON DELETE SET NULL,
    CONSTRAINT fk_currency FOREIGN KEY (currency) REFERENCES currency(id) ON DELETE SET NULL
);

