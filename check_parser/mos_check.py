import os

os.execlp("psql", "-d", "expenses3", "-c", "select name, address, id from market_or_supplier", "--tuples-only", "--no-align")
