-- BASE MOVIER SERVICE DATA
CREATE TABLE IF NOT EXISTS country (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE,
    code VARCHAR(5) UNIQUE
);

CREATE TABLE IF NOT EXISTS studio (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50),
    country_id BIGINT REFERENCES country
);

CREATE TABLE IF NOT EXISTS movies (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50),
    date_of_release DATE,
    poster_url TEXT,
    rating VARCHAR,
    studio_id BIGINT REFERENCES studio
);

CREATE TABLE IF NOT EXISTS persons (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50),
    surname VARCHAR(50),
    second_name VARCHAR(50),
    date_of_birth DATE,
    profile_picture_url TEXT,
    CONSTRAINT person_unique UNIQUE (name, surname, second_name, date_of_birth)
);

CREATE TABLE IF NOT EXISTS external_info (
    movie_id BIGINT REFERENCES movies,
    ext_id VARCHAR(50),
    ext_source VARCHAR(50),
    ext_rating DECIMAL,
    rates BIGINT,
    CONSTRAINT external_info_unique UNIQUE (movie_id, ext_id)
);

CREATE TABLE IF NOT EXISTS genres (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE
);

CREATE TABLE IF NOT EXISTS movie_person (
    movie_id BIGINT NOT NULL REFERENCES movies,
    role VARCHAR(50),
    person_id BIGINT NOT NULL REFERENCES persons,
    CONSTRAINT movie_person_unique UNIQUE (movie_id, role, person_id)
);

CREATE TABLE IF NOT EXISTS movie_genre (
    movie_id BIGINT NOT NULL REFERENCES movies,
    genre_id BIGINT NOT NULL REFERENCES genres,
    CONSTRAINT movie_genre_unique UNIQUE (movie_id, genre_id)
);

-- USER SERVICE DATA
CREATE TABLE IF NOT EXISTS "user" (
    id BIGSERIAL PRIMARY KEY,
    tg_id BIGINT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS user_rating (
    user_id BIGINT REFERENCES "user",
    movie_id BIGINT REFERENCES movies,
    rating SMALLINT,
    CONSTRAINT user_rating_unique UNIQUE (movie_id, user_id)
);