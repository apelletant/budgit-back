Simple go logger

***Example usage:*** <br />
logger := NewCustomLogger(true) // enable debug logs <br />
logger.Info("Application started", "port", 8080) <br />
logger.Error("Database connection failed", err, "host", "localhost") <br />
logger.Warn("Deprecated API endpoint used", "endpoint", "/old-api") <br />
logger.Debug("User data", "userID", 123, "email", "user@example.com") <br />
logger.Success("User created successfully", "userID", 456) <br />
