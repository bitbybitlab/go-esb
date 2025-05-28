INSERT INTO public.users(username, password, first_name, last_name) VALUES('ivanmolodec', '123', 'Ivan', 'Smirnov');
INSERT INTO public.users(username, password, first_name, last_name) VALUES('denisushakov', '456', 'Denis', 'Ushakov');
INSERT INTO public.users(username, password, first_name, last_name, middle_name, email)
    VALUES('thvvmas', '789', 'Vladislav', 'Permichev', 'Vladimirovich', 'blkycc1x@gmail.com');

INSERT INTO public.systems(name) VALUES('1C');

INSERT INTO public.connection_authentications(
	name, system, username, password, token)
	VALUES ('Подключение к 1С',
	'3072fb9b-5da7-45c4-9458-256336b04956',
	'test', 'test', '');