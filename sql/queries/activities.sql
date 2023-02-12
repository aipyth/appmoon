-- name: AddActivity :exec
insert into activities
(class_name, title, timestamp)
values ($1, $2, now());
