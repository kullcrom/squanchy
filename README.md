# squanchy

users table:
```
id - INT
email - VARCHAR(50)
password - VARCHAR(60)
first_name - VARCHAR(25)
last_name - VARCHAR(25)
PRIMARY KEY - (id)
```

properties table:
```
id - INT
FOREIGN KEY - user_id - INT
street - VARCHAR(25)
city - VARCHAR(20)
state - VARCHAR(2)
purchase_price - FLOAT
exit_strategy - SMALLINT
```

Can we calculate taxes and insurance based on industry and state rates?

(id INT unsigned NOT NULL AUTO_INCREMENT, email VARCHAR(50) NOT NULL, password VARCHAR(60) NOT NULL, first_name VARCHAR(25) NOT NULL, last_name VARCHAR(25) NOT NULL, PRIMARY KEY (id));

(id INT unsigned NOT NULL AUTO_INCREMENT, user_id INT unsigned, INDEX user_ind (user_id), FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE, street VARCHAR(25) NOT NULL, city VARCHAR(20) NOT NULL, state VARCHAR(2) NOT NULL, purchase_price FLOAT NOT NULL, exit_strategy SMALLINT NOT NULL, PRIMARY KEY (id))
