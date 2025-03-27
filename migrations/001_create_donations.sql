CREATE TABLE donations (
    id SERIAL PRIMARY KEY,
    amount INTEGER NOT NULL,
    status VARCHAR(20) NOT NULL,
    authority VARCHAR(50) UNIQUE,
    ref_id BIGINT DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);