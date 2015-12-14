package utils;

import java.time.Duration;
import java.time.Instant;

/**
 * Created by Andre on 05-12-2015.
 */
public class Crono {
    private Instant initialInstant;
    private Instant finalInstant;
    private boolean isStopped;

    public Crono() {
        isStopped = true;
    }

    public void start() {
        isStopped = false;
        initialInstant = Instant.now();
    }

    public Duration stop() {
        this.finalInstant = Instant.now();
        this.isStopped = true;
        return Duration.between(initialInstant, finalInstant);
    }

    public Duration getElapsedTime() {
        return Duration.between(initialInstant, (this.isStopped ? Instant.now() : finalInstant));
    }
}
