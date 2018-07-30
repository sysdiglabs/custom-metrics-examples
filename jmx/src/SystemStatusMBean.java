public interface SystemStatusMBean {
    Integer getNumberOfSecondsRunning();
    String getProgramName();
    Long getNumberOfUnixSecondsRunning();
    Boolean getSwitchStatus();
}
