# Gomodoro

Gomodoro is a Pomodoro timer written in Golang that additionally exports Pomodoro session data as metrics for ingestion into analytics and monitoring systems.

This allows users to track their own works habits over time as well as leverage the capabilities or monitoring systems to track, and even dispatch alerts on, the number of pomodoro sessions completed over time.

## Use

Gomodoro can be used to set timers such as

`./Gomodoro -T 3600`

or

`./Gomodoro --time 3600`
