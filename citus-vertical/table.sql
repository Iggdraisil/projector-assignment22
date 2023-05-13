CREATE TABLE books (
    id integer NOT NULL,
    title text,
    author text,
    year integer,
);
select create_distributed_table('books', 'id');
