INSERT INTO agendamentos (propriedade_id, dia_agendamento, checkout) VALUES (1, STR_TO_DATE('02-01-2023', '%d-%m-%Y'), '10:10:10');


SELECT * FROM agendamentos a where X


select a.propriedade_id, a.dia_agendamento, a.checkin, a.checkout, a.observacoes from agendamentos a INNER JOIN propriedades p ON p.id = a.propriedade_id inner join usuarios u on u.id = p.cliente_id where u.id = 1;

