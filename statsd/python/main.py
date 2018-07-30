import statsd
import time

if __name__ == '__main__':

    # Create a new connection for the client
    connection = statsd.Connection(
        host='localhost',
        port=8125,
        sample_rate=1,
    )

    # Create the client
    client = statsd.Client("python", connection)

    # Create counter
    counter = client.get_counter("counter")

    # Create gauge
    gauge = client.get_gauge("gauge")

    # Create average
    average = client.get_average("average")


    while True:
        # Create a timer
        timer = client.get_timer("timer")

        # Will send the elapsed time once all the block has been executed
        with timer:
            counter += 1  # Increment by one the counter

            gauge.set('foo', 10) # Send a gauge of 10

            average.send('avg', 5)

        time.sleep(1)

