CREATE TABLE IF NOT EXISTS mailing_messages(
   mailing_id INT NOT NULL,
   email VARCHAR (300) NOT NULL,
   title VARCHAR (200) NOT NULL,
   content TEXT NOT NULL, 
   insert_time TIMESTAMP NOT NULL
);
