
#!/bin/bash
{
    cat <<-eof
    DROP DATABASE IF EXISTS blog;
    CREATE DATABASE blog;
    \c blog
eof
    for i in $(ls -1 ./*.sql | sort)
    do
        echo "\\echo 'running $i'"
        cat "$i"
    done
    cat <<-eof
    SELECT pg_terminate_backend(pg_stat_activity.pid) 
    FROM pg_stat_activity 
    WHERE pg_stat_activity.datname = 'blog' AND pid <> pg_backend_pid(); 
eof
} | psql -v ON_ERROR_STOP=1 --echo-errors -U gitpod --host localhost -p 5432