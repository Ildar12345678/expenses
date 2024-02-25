import json
import os
import sys


map_subcategory = {
    1: "лекарство", 2: "обследования", 3: "уход за собой", 4: "профилактика",  # здоровье

    5: "одежда", 6: "обувь", 7: "аксессуары", 8: "мебель", 9: "электроника", 10: "хозтовары",
    11: "для ремонта", 12: "книги", 13: "спорт", 14: "развлечения", 15: "косметика", 16: "другое",  # непродукты

    17: "для дома", 18: "фр/сфр/ор", 19: "неполезное", 20: " не дома", 35: "животным",  # продукты

    21: "межгород", 22: "такси", 23: "общественный",  # проезд

    24: "кафе", 25: "культуры", 26: "заказ еды", 27: "путешествия", 28: "другое",  # развлеченья

    29: "услуги", 30: "государству", 31: "подарки", 32: "сотовый", 33: "благотвор", 34: "другое",  # другое
}

# читаем ключи подкатегорий из файла, в который они попадают из echo '...' > subcats
with open("subcats") as f:
    string = f.read().removesuffix('\n')
    order_subcategory = list(map(lambda s: int(s), string.split(" ")))

# если такого продавца еще нет, то нужно добавить
# этот параметр будет перезадан: делаю запрос в базу на проверку наличия такого продавца
# значение из json проверяю по этому списку
to_add_mos = True

# Check if the user provided a file path argument

if len(sys.argv) != 3:
    print("Usage: python my_script.py file_path")
    sys.exit(1)

# Get the file path from the command-line arguments
json_file_path = sys.argv[2]
# Open the JSON file
with open(json_file_path, "r") as f:
    # Load the JSON data from the file
    json_data_from_check = json.load(f)

    data_time = str(json_data_from_check[0]["ticket"]["document"]["receipt"]["dateTime"]).lower().split('t')[0]
    retail_place = str(json_data_from_check[0]["ticket"]["document"]["receipt"]["retailPlace"]).lower()
    retail_place_address = str(json_data_from_check[0]["ticket"]["document"]["receipt"]["retailPlaceAddress"]).lower()
    products = json_data_from_check[0]["ticket"]["document"]["receipt"]["items"]


# назначаем подкатегории
for i, product in enumerate(products):
    product["subcat"] = order_subcategory[i]

for i in range(len(products)):
    # print("i", products[i])
    for j in range(i+1, len(products)):
        # print("j", products[j])
        if products[i]['name'] == products[j]['name']:
            products[i]['quantity'] += 1
            continue
lst_products = {}
# list(products).reverse()

for i in range(len(products)-1,-1,-1):
    lst_products[hash(f"{products[i]['name']}")] = products[i]


with open("from_expense") as f:
    expense_already_exist = f.read().split('\n')

map_expense_already_exist = {}
for s in expense_already_exist:
    if s == '':
        continue
    arr = s.split("|")
    map_expense_already_exist[arr[0]] = arr[1]

# print(map_expense_already_exist)
# если в базе такой товар уже есть, то название товара мы подменяем на его id,
# чтобы потом вставить сразу в чек
for k, v in map_expense_already_exist.items():
    for _,product in lst_products.items():
        if product["name"] == v:
            product["name"] = k


# удаляем ненужные ключи
for _, product in lst_products.items():
    product['nds'] = 10 if product['nds'] == 2 else 20
    del product['productType']
    del product['paymentType']
    del product['sum']
    try:
        del product['productCodeDataError']
    except KeyError:
        continue

# for k, v in lst_products.items():
#     print(k, v)
# #
# sys.exit()
map_city = {1: "снежинск", 2: "челябинск", 3: "екатеринбург", 4: "казань", }
city = ""
# задаем город покупки
for k, v in map_city.items():
    split_word = "область"
    if split_word in retail_place_address:
        arr = retail_place_address.split(split_word)
        retail_place_address = arr[1]
    if v in retail_place_address:
        city = v
        break
    else:
        city = map_city[int(sys.argv[1])]
        retail_place_address = city
# ***************************************************************
# в этом блоке проверяю, есть ли уже такой продавец
# проверка идет по названию и адресу (arr[0], arr[1])
# еще запрашиваю из базы id, который будет подставлен в запрос, если такой продавец уже есть
# to_add_mos выставляю в False, если такой продавец уже есть
# информацию беру из файла mos_check, в который записываю информацию из базы
with open("mos_check") as f:
    mos_already_exist = f.read().split('\n')

map_mos_already_exist = {}
for s in mos_already_exist:
    if s == '':
        continue
    arr = s.split("|")
    map_mos_already_exist[arr[2]] = arr[0] + arr[1]

mos_id = None
# print(map_mos_already_exist)
for k, v in map_mos_already_exist.items():
    if v == retail_place + retail_place_address:
        to_add_mos = False
        mos_id = k
# ***************************************************************
# здесь проверяю равенство двух списков.
# они должны быть равны, иначе подкатегория у товара будет неверная
if len(products) != len(order_subcategory):
    os.system("echo 'look at 112 string of main.py'")
    sys.exit(1)
# ***************************************************************
# здесь проверяю равенство двух списков.

# *********************
# НАЧИНАЕМ ТРАНЗАКЦИЮ
# *********************
print("begin;")

if to_add_mos:
    # *********************
    # ВСТАВЛЯЕМ ПРОДАВЦА
    # *********************
    string_insert_mos = f"insert into market_or_supplier (name, address) values ('{retail_place}', '{retail_place_address}');"
    print(string_insert_mos)
    # *********************
    # ВСТАВЛЯЕМ ИНФОРМАЦИЮ О ПОКУПКЕ
    # *********************
    string_insert_purchase = f"insert into purchase (purchase_date, city_id, online, description, mos_id) " \
                             f"select '{data_time}', city.id, 'false', 'some description', mos.last_value " \
                             f"from city, market_or_supplier_id_seq as mos " \
                             f"where city.city = '{city}';"
else:
    # *********************
    # ВСТАВЛЯЕМ ИНФОРМАЦИЮ О ПОКУПКЕ ЕСЛИ ТАКОЙ ПРОДАВЕЙ УЖЕ ЕСТЬ
    # *********************
    string_insert_purchase = f"insert into purchase (purchase_date, city_id, online, description, mos_id) " \
                             f"select '{data_time}', city.id, 'false', 'some description', {mos_id} " \
                             f"from city where city.city = '{city}';"
print(string_insert_purchase)

for _,product in lst_products.items():
    price = round(product['price'] / 100)
    quantity = product['quantity']
    string_insert_expense = string_insert_purchase_check = ""
    if not str(product['name']).isdigit():
        # *********************
        # ВСТАВЛЯЕМ ИНФОРМАЦИЮ О ТОВАРАХ
        # *********************
        string_insert_expense = f"insert into expense (name, subcat_id, nds) values " \
                                f"('{product['name']}', {product['subcat']}, {product['nds']});"
        # *********************
        # ВСТАВЛЯЕМ ИНФОРМАЦИЮ О ЧЕКЕ ДЛЯ НОВЫХ ТОВАРАХ
        # *********************
        string_insert_purchase_check = f"insert into purchase_check (expense_id, purchase_id, count, price) " \
                                       f"select e.last_value, p.last_value, {quantity}, {price} from expense_id_seq as e, purchase_id_seq as p;"

    else:
        # *********************
        # ВСТАВЛЯЕМ ИНФОРМАЦИЮ О ЧЕКЕ ДЛЯ СУЩЕСТВУЮЩИХ ТОВАРАХ
        # *********************
        string_insert_purchase_check = f"insert into purchase_check (expense_id, purchase_id, count, price) " \
                                       f"select {product['name']}, p.last_value, {quantity}, {price} from purchase_id_seq as p;"

    print(string_insert_expense)
    print(string_insert_purchase_check)
# *********************
# КОММИТ ТРАНЗАКЦИИ
# *********************
print("commit;")
