Задача:

  In memory simple DB
  В процессе реализации сервиса проверки данных возникла необходимость в организации in memory кэша с возможностью быстрого поиска по разным полям.
  Структура данных представлена следующим набором полей:
  "account": "234678", //long
  "name": "Иванов Иван Иванович", //string
  "value": "2035.34" //double
  Количество записей заранее не определено и может меняться динамически.
  ВОПРОС: Необходимо организовать хранение этих записей в памяти с соблюдением требований:
  1. предоставить возможность добавлять новые записи;
  2. предоставить возможность удалять более не нужные записи;
  3. предоставить возможность изменять запись;
  4. получать полный набор записи по любому из полей с одинаковой алгоритмической сложностью (не медленнее
  log(n));
  5. выбрать наиболее экономный способ хранения данных в памяти.
  Важно: Нужно обосновать выбор структур данных и алгоритмов относительно требований
  Реализовать это в виде микросервиса в конвенции REST-APi 

Результат:
  
  1.Функция запроса Post /users куда так же необходимо передать запрос в формате Json с параметрами по указанной структуре
  
  2.Функция запроса DELETE /users/:Account где Account является необходимым для удаления Account
  
  3.Функция запроса PUT /users/:Account где Account является необходимым для изменения Account, куда так же необходимо передать запрос в формате Json с параметрами по указанной структуре
  
  4.Функция запроса Get /users  и возможность указания  одного фильтра в формате ?account=XXX; ?name=XXX; ?value=XXX.
Система поиска реализована через  хешмапу (O(n)), для формирования который использовались ссылки на объекты и, следовательно, не привели к заметному изменению использованной памяти
  
  5. В памяти вся информация хранится внутри таблицы и так же на основе указателей создаются 3 хашмапы для реализации поиска, тк в рамках задания нельзя было использовать стороннее хранилище и  мемхешь не подходит по используемой памяти  был выбран максимально  просто и в то же время экономных вариант.
  
Выбор создания дополнительных 3 хешмап основан на том что формирование b3 получилось бы долгим и сложным, а разницы в рамках данного проекта не будет. А открытый поиск по хешмапе это самый быстрый способ поиска который я придумал реализовать  исоздание нового объекта в хешмапах значительно проще чем создание нового элемента и индекса с последующий сортировкой по другим методам. 

Описание Swagger бедут добавлено чуть позже.
