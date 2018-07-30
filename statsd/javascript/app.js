const SDC = require('statsd-client');

const sdc = new SDC({host: '127.0.0.1'});


setInterval(() => {
    const timer = new Date();

    // Increment counter by one.
    sdc.increment('node.counter');

    // Increment counter by 10
    sdc.increment('node.counter_10', 10);

    // Set gauge to 10
    sdc.gauge('node.gauge', 10);

    // Calculates time diff of time between the variable and
    // when the function was called
    sdc.timing('node.timer', timer);

    // Set will count just 2 elements since '50' is repeated
    sdc.set('node.set', 50);
    sdc.set('node.set', 100);
    sdc.set('node.set', 50);

    // Histogram with tags
    sdc.histogram('node.histogram', 10, {foo: 'bar'});
}, 1000);

