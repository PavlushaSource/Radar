# Dogs Radar

<img width="100" alt="logo" align="right" src="https://github.com/user-attachments/assets/1e257c78-76a6-41dc-b34d-d898cf4489e8">

Dogs Radar - инсутремент для симуляции взаимодействия объектов на плоскости

Для увеличения красоты и понимания происходящего, в качестве объектов в приложении используются собачки

## Правила взаимодействия собак

Взаимодействие собак между собой определяется по следующим правилам:

1. Если две собаки находятся на расстоянии не превышающем `R0`, то они пытаются начать драку с вероятностью 1

2. Если две собаки находятся на расстоянии `R1`, так, что `R1 > R0`, они начинают рычать с вероятностью обратно
   пропорциональной квадрату расстояния между ними

3. Если вокруг собаки нет соперников, она перемещается согласно текущему правилу

## Функциональность приложения

1. Обработка на движке более чем 200 тысяч собак за менее 500 мс
2. Подсветка взаимодействия собак
3. Возможность регулирования следующих параметров симуляции:
    - Количество собак
    - Скорость обновления перемещений и состояний собак
    - Радиусы `R0` и `R1`, задающие ограничения максимального расстояния драки и рычания соответственно
    - Паттерн движения собак
        * Свободное (Simple)
        * Движение прямо с отклонением на некоторый угол (Vector)
    - Функция пересчета расстояний
        * Евклидово расстояние
        * Манхэттенское расстояние
        * Половина длины окружности, построенной через 2 рассматриваемых собаки
4. Возможность зума и перемещения камеры по карте для детального наблюдения за собаками
5. Возможность выбора темы и возможность вернуться на главную страницу для перезапуска взаимодействия собак с вновь
   введёнными параметрами

## Недоработанная функциональность

1. Возможность поддержки барьеров (на демонстрации показана возможность движка корректно обрабатывать случаи
   столкновения собак с барьерами)
2. Зум может быть плавнее, плывёт представление границы и собаки выходят немного за границу
3. Инструмент [Fyne](https://github.com/fyne-io/fyne) не оправдал ожиданий по производительности, подробнее сравнение
   представлено
   в [презентации](https://docs.google.com/presentation/d/1uTHLYBgTpydVI3z8IItD46WBHYIrUnuaMW_fcr7SpP4/edit?usp=sharing)

## Демонстрация работы приложения

Will be soon...

## Запуск приложения

Для запуска нашего приложения вы можете использовать готовые сборки под разные платформы,
представленные [здесь](https://github.com/PavlushaSource/Radar/releases)

Для самостоятельной сборки и запуска приложения используйте:

```shell 
  make run
```

Самостоятельно можно посмотреть бенчмарки с профайлером для движка отдельно, для этого используйте:
```shell
   make run-engine
   go tool pprof -http=:8080 cpu.pprof
```

Затем нужно зайти на localhost:8080 для просмотра подробного дерева профайлинга или флэйм графа.

## Архитектура и используемые решения

### [Описание алгоритма](algo.md)

### [Инструменты и архитектурная документация](arch.md)

