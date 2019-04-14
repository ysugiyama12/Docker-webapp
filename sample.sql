CREATE TABLE my_user (
    id bigserial primary key,
    name varchar(255) NOT NULL,
    email varchar(255),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);

create function set_update_time() returns opaque as '
  begin
    new.updated_at := ''now'';
    return new;
  end;
' language 'plpgsql';

create trigger update_tri before update on my_user for each row
  execute procedure set_update_time();