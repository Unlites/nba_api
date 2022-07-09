CREATE TABLE teams
(
    id serial NOT NULL unique,
    name varchar(255) NOT NULL
);

CREATE TABLE players
(
    id serial NOT NULL unique,
    name varchar(255) NOT NULL,
    team_id int NOT NULL
);

CREATE TABLE games
(
    id serial NOT NULL unique,
    home_team_id int NOT NULL,
    visitor_team_id int NOT NULL,
    score varchar(255) NOT NULL,
    won_team_id int NOT NULL
);

CREATE TABLE stats
(
    id serial NOT NULL unique,
    game_id int references games (id) on delete cascade NOT NULL,
    player_id int NOT NULL,
    points int NOT NULL,
    rebounds int NOT NULL,
    assists int NOT NULL
);