# Вопросы на вакансию golang
- шаблоны/паттерны проектирования
- [выбор монолит/микросервисы](https://habr.com/ru/post/459810/)
- что происходит при вводе урла в браузере
- индексы, детали внутреннего устройства
    - [в pgsql](https://habr.com/ru/company/postgrespro/blog/330544/)
    - [wiki](https://ru.wikipedia.org/wiki/%D0%98%D0%BD%D0%B4%D0%B5%D0%BA%D1%81_(%D0%B1%D0%B0%D0%B7%D1%8B_%D0%B4%D0%B0%D0%BD%D0%BD%D1%8B%D1%85))
    - [еще куча всего](https://ylianova.ru/raznoe-2/mysql-indeksy-indeksy-v-mysql.html)
- [отличия http2 от http1](https://www.digitalocean.com/community/tutorials/http-1-1-vs-http-2-what-s-the-difference)
- устройства рантайма голанга
    - [рантайм это](https://golang.org/doc/faq#runtime)
- многопоточность, [на уровне системных вызовов](https://medium.com/a-journey-with-go/go-goroutine-os-thread-and-cpu-management-2f5a5eaf518a)
- Про асинхронность и горутин
- Разница в запросах хед, гет, пост и прочих
- [Асинхронное программирование с примерами](https://proglib.io/p/asynchrony)
- [habr c-вызовы в го: принципы и производительноcть](https://habr.com/ru/company/intel/blog/275709/)
- [Как работает гарбэдж коллектор](https://ru.wikipedia.org/wiki/%D0%A1%D0%B1%D0%BE%D1%80%D0%BA%D0%B0_%D0%BC%D1%83%D1%81%D0%BE%D1%80%D0%B0)
- [гарбэдж коллектор в го](https://www.youtube.com/watch?v=CX4GSErFenI)
- Дизайн сервисов - выбор индексов в базе, где встроить мэссэдж-брокер
- [Композитный индекс как устроен внутри](https://habr.com/ru/post/247373/)
- [реализации кеша в http](https://zametkinapolyah.ru/servera-i-protokoly/tema-13-keshirovanie-v-http-mexanizmy-klientskogo-i-servernogo-kesha-v-http.html)
- какие сервисы нужны для mvp по продукту
- как спроектировать БД -безопасность при работе с БД.
- Чем отличается слайс от массива
- С какими фреймворками в go работал
- Как ограничить кол-во подключение к базе
- Что делает данный кусок кода (сортировка)
- Какие сортировки используете?
- массивы слайсы интерфейсы, синхронизация многопоточности,
- как устроен Go "под капотом", [кап теорема](https://ru.wikipedia.org/wiki/%D0%A2%D0%B5%D0%BE%D1%80%D0%B5%D0%BC%D0%B0_CAP)
- [почему бинарник go весит значительно больше сишного?](https://stackoverflow.com/a/40141212)
- [Как работает рантайм с рутинами в разрезе тредов ОС?](https://medium.com/a-journey-with-go/go-goroutine-os-thread-and-cpu-management-2f5a5eaf518a)
- С какими базами данных работал?
- Как будешь разбираться с проблемой медленного запроса в реляционной БД?
- Что будет при чтении из nil канала - поток заблокируется навсегда
- Что будет при записи в  nil канал - поток заблокируется навсегда
- Как работает шедулер горутин [1часть](https://habr.com/ru/post/478168) / [2часть](https://habr.com/ru/post/489862/)
- Как оптимизировать select в postgres
- про опыт и особые достижения в каждой из организаций
- откуда берешь знания и как учишься чему-то новому
- как относишься к большим задачам и как их декомпозируешь
- готов ли работать в стартапе  \ энтерпрайзе
- в какую сторону готов развиваться
- хочешь ли в devops
- что лучше перфекционизм или mvp (обрати внимание, тут надо быть объективным)


Доп:
- что такое [binding](https://stackoverflow.com/questions/27014955/socket-connect-vs-bind), модель OSI, маршрутизация
- как будешь вносить изменения в базу данных (знаю, что не нужно отвечать, что каждый раз будешь переписывать адаптер)
- как работает [keep-alive](https://www.imperva.com/learn/performance/http-keep-alive/#:~:text=HTTP%20keep%2Dalive%2C%20a.k.a.%2C,connections%20close%20after%20each%20request.&text=Enabling%20the%20keep%2Dalive%20header,resources%20over%20a%20single%20connection.)
- http
- dns
- как работают [каналы в go](https://medium.com/rungo/anatomy-of-channels-in-go-concurrency-in-go-1ec336086adb#:~:text=What%20are%20the%20channels%3F,data%20from%20the%20same%20channel.)
- QUIC
- откуда в операционной системе лимит на количество открытых соединений

Надо уметь:
- написать [паттерн Singleton](https://kodazm.ru/articles/singlton-pattern-na-go/#.YB8TLhMzZYg)
- понимать, какие объекты необходимо защищать при конкурентном доступе
- отстаивать свое мнение

Админство:
- linux
- ядро
- iptables
- bsd
- windows
- сети
- K8s
- docker
- docker-compose