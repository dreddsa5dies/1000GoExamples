#!/bin/bash
# создание директорий в моем формате

for DIR in {1..20}; do
if [ ! -d "$DIR" ]; then
	if [ "$DIR" -lt 10 ]; then
		mkdir '00'$DIR
	else mkdir '0'$DIR
	fi
fi
done
exit 0
