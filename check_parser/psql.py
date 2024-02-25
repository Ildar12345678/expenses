import os

os.execlp("psql", "-d", "expenses3", "-c", "select id, name from expense", "--tuples-only", "--no-align")
