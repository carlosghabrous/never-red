CREATE TABLE IF NOT EXISTS movements (
    id SERIAL PRIMARY KEY,
    mov_date DATE NOT NULL,
    mov_type VARCHAR(10), 
    commerce VARCHAR(50),
    category VARCHAR(50),
    currency VARCHAR(3),
    amount FLOAT, 

    CONSTRAINT fk_movtype(mov_type) REFERENCES mov_type(id) ON DELETE SET NULL,
    CONSTRAINT fk_category(category) REFERENCES category(id) ON DELETE SET NULL,
    CONSTRAINT fk_currency(currency) REFERENCES currency(id) ON DELETE SET NULL
);

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
