--
create table items (
    id serial primary key not null,
    name varchar(256) not null,
    code varchar(256) not null,
    unit varchar(256) not null,
    cost numeric(13, 5) not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);
--
CREATE FUNCTION set_update_time() RETURNS OPAQUE AS '
  begin
    new.updated_at := ''now'';
    return new;
  end;
' LANGUAGE plpgsql;
--
CREATE TRIGGER update_tri BEFORE UPDATE ON items
FOR EACH ROW EXECUTE PROCEDURE set_update_time();
