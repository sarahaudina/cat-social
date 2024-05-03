CREATE TABLE cats(
    cat_id serial primary key,
    user_id serial,
    name VARCHAR(255) NOT NULL,
    race VARCHAR(255) NOT NULL,
    age INTEGER NOT NULL,
    CONSTRAINT fk_user
      FOREIGN KEY(user_id) 
        REFERENCES user(user_id)
);