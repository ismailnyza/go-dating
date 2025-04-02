package configuration

import "github.com/go-sql-driver/mysql"


type DbConfig struct {
    Host string
    Port string
    User string
    Password string
    DBName string
    TimeZone string
}

//similar to a constructor in other languages
func NewDbConfig() *DbConfig {
    return &DbConfig{
    Name: "root",
}

func (config *DbConfig)EstablishConfig() error{
    mysqlConfig := mysql.Config{
        User: config.User,
        Passwd: config.Password,
        Net: "tcp",
        Addr: config.Host + ":" + config.Port,
        DBName: config.DBName,
        TimeZone: config.TimeZone,
    }

    db , err :+ sql.Open("mysql", mysqlConfig.FormatDSN())
    if err != nil {
        return fmt.Errorf("Error opening database: %v", err)
    }

    if err := db.Ping(); err != nil {
        return fmt.Errorf("Error pinging database: %v", err)
    }

    config.DB = db
    fmt.Println("Database connection established to %s", config.DBName)

}

func (config *DbConfig)CloseConnection(){
    if config.DB == nil {
    config.DB.Close()
    fmt.Println("Database connection closed")
    }
}



func main(){
    DbConfig := DbConfig{
        Host: "localhost",
        Port: "3306",
        User: string,
        Password: string
    }

    if err := DbConfig.EstablishConfig(); err != nil {
        fmt.Println("Error establishing database connection: %v", err)
    }

    defer DbConfig.closeConnection()
}
