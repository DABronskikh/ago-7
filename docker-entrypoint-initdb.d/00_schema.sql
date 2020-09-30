CREATE TABLE flights
(
    id              BIGSERIAL PRIMARY KEY,
    cost            BIGSERIAL,
    fromIATA        BIGSERIAL,
    toIATA          BIGSERIAL,
    timeTravel      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    timeDeparture   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
