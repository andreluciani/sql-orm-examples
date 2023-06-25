# Create database, models and add seed data used in the definitions of database concepts
psql -U gitpod -d blog -a -f ./create-databases.sql
psql -U gitpod -d blog -a -f ./blog-models.sql
psql -U gitpod -d blog -a -f ./blog-seed.sql