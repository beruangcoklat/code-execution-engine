drop keyspace if exists code_execution_engine;
create keyspace code_execution_engine with replication = {'class' : 'SimpleStrategy', 'replication_factor':3};
use code_execution_engine;

drop type if exists testcase;
create type testcase (
    input text,
    expected_output text
);

drop table if exists problem;
create table problem (
   id uuid primary key,
   title text,
   description text,
   testcases list<frozen <testcase>>
);

drop table if exists language;
create table language (
    id int primary key,
    name text
);
insert into language (id, name) values (1, 'C++');
insert into language (id, name) values (2, 'Java');

