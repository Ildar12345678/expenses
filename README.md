Через expenses.sql можно создать базу со всеми таблицами. Нужен PostgreSQL
Делаем:
1. go build main.go
2. ./main - запускаем сервер
3. cd web/first-react-tutorial/
4. npm start - запускаем веб-клиента. Теперь можно ручками вставлять расходы
Чтобы парсить чек, делаем:
1. cd check_parser
2. echo '1 2 3' > subcats. 1,2,3 - подкатегории из map_subcats
3. python3 final.py NUM file.json. NUM - номер города из main.py (102 строка), file.json - чек
