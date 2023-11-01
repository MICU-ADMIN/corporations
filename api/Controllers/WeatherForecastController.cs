// Import the necessary namespace for ASP.NET Core MVC controllers.
using Microsoft.AspNetCore.Mvc;

// Define a controller class named "WeatherForecastController" within the "api.Controllers" namespace.
namespace api.Controllers
{
    // Indicate that this class is an API controller.
    [ApiController]

    // Specify the route for this controller, which in this case is based on the controller name.
    [Route("[controller]")]
    public class WeatherForecastController : ControllerBase
    {
        // Define an array of weather summaries.
        private static readonly string[] Summaries = new[]
        {
            "Freezing", "Bracing", "Chilly", "Cool", "Mild", "Warm", "Balmy", "Hot", "Sweltering", "Scorching"
        };

        // Create a logger instance for this controller.
        private readonly ILogger<WeatherForecastController> _logger;

        // Constructor to initialize the controller with a logger.
        public WeatherForecastController(ILogger<WeatherForecastController> logger)
        {
            _logger = logger;
        }

        // Define an HTTP GET endpoint that returns a list of weather forecasts.
        [HttpGet(Name = "GetWeatherForecast")]
        public IEnumerable<WeatherForecast> Get()
        {
            // Generate a list of weather forecasts using random data.
            return Enumerable.Range(1, 5).Select(index => new WeatherForecast
            {
                Date = DateOnly.FromDateTime(DateTime.Now.AddDays(index)),
                TemperatureC = Random.Shared.Next(-20, 55),
                Summary = Summaries[Random.Shared.Next(Summaries.Length)]
            })
            .ToArray();
        }
    }
}
