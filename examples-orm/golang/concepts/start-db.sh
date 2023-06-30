
#!/bin/bash
{
cat <<-eof
    DROP DATABASE IF EXISTS quotes_db;
    CREATE DATABASE quotes_db;
    \c quotes_db
eof
    cat "./quotes.sql"
} | psql -U gitpod --host localhost -p 5432