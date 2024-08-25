-- For creating the tables in db
DROP TABLE IF EXISTS "Users";
DROP TABLE IF EXISTS Trips;
-- DROP TABLE IF EXISTS TripLists;
DROP TABLE IF EXISTS TripItems;

CREATE TABLE "Users" (
    user_id SERIAL PRIMARY KEY,
    email TEXT NOT NULL,
    username VARCHAR(55) NOT NULL,
    password VARCHAR(255) NOT NULL
);


CREATE TABLE Trips (
    trip_id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- CREATE TABLE TripLists (
--     trip_list_id SERIAL PRIMARY KEY,
--     name TEXT NOT NULL,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     user_id INT NOT NULL REFERENCES "User" ON DELETE CASCADE,
--     trip_id INT NOT NULL REFERENCES Trip ON DELETE CASCADE
-- );

CREATE TABLE TripItems (
    trip_item_id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    trip_list_id INT NOT NULL REFERENCES TripList ON DELETE CASCADE,
    is_packed BOOLEAN DEFAULT FALSE
);

