// Define a class named "WeatherForecast" within the "api" namespace.
namespace api
{
    // This class represents weather forecast data.

    public class WeatherForecast
    {
        // Property to store the date of the weather forecast.
        public DateOnly Date { get; set; }

        // Property to store temperature in degrees Celsius.
        public int TemperatureC { get; set; }

        // Property to calculate temperature in degrees Fahrenheit based on TemperatureC.
        public int TemperatureF => 32 + (int)(TemperatureC / 0.5556);

        // Property to store a summary of the weather forecast (nullable string).
        public string? Summary { get; set; }
    }
}
