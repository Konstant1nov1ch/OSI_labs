# Бенчмарки для 12th Gen Intel Core i7-1255U

Этот репозиторий содержит скрипты на Go для запуска бенчмарков с помощью утилиты `stress-ng` на системе с процессором 12th Gen Intel Core i7-1255U и 16 GB оперативной памяти.

## Система
- Процессор: 12th Gen Intel Core i7-1255U (12 ядер)
- Оперативная память: 16 GB

## Инструкции по запуску

Для запуска бенчмарков необходимо установить `stress-ng` на вашей системе. Следуйте инструкциям по установке stress-ng для вашего дистрибутива Linux.

### вариант для бенчмарков
- cpu: [ipv4checksum,gray]; 
- cache: [cache-prefetch,l1cache]; 
- io: [iomix,ioport]; 
- memory: [mmaphuge-mmaps,misaligned-method]; 
- network: [sockdiag,netlink-task]; 
- pipe: [pipeherd-yield,pipe-data-size]; 
- sched: [schedpolicy,yield]

