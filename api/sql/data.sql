SELECT * FROM agendamentos a where X


select a.propriedade_id, a.dia_agendamento, a.checkin, a.checkout, a.observacoes from agendamentos a INNER JOIN propriedades p ON p.id = a.propriedade_id inner join usuarios u on u.id = p.cliente_id where u.id = 1;

//DATAS

select * from agendamentos where extract(year from dia_agendamento) = 2023;

select * from agendamentos where extract(month from dia_agendamento) = 2;

select * from agendamentos where extract(day from dia_agendamento) = 40;

select * from agendamentos where dia_agendamento = DATE "2023-01-31";


select * from agendamentos where extract(year from dia_agendamento) = 2023 order by dia_agendamento asc;

select * from agendamentos where extract(year from dia_agendamento) = 2023 AND extract(month from
dia_agendamento) = 10 order by dia_agendamento asc;


select * from agendamentos where extract(year from dia_agendamento) = 2023 AND extract(month from
dia_agendamento) = 10 and extract(day from dia_agendamento) = 2 order by dia_agendamento asc;