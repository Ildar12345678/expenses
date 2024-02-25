import sys
from os import execlp
import os
# это аргумент, в котором хранится название файла-json
city_id = sys.argv[1]
json = sys.argv[2]
os.system("python3 mos_check.py > mos_check")
os.system("python3 psql.py > from_expense")
os.system(f"python3 main.py  {city_id} {json} > file.sql")
os.system("psql -d expenses3 -f file.sql")

# execlp("python3", "psql.py", ">", "from_expense")
# execlp("python3", "main.py", "1.json", ">", "file.sql")
# execlp("psql", "-d", "expenses3", "-f", "file.sql")
# "python3 psql.py > from_expense"
# "python3 main.py 1.json > file.sql"
# "psql -d expenses3 -f file.sql"