// Create a web application builder using command-line arguments.
var builder = WebApplication.CreateBuilder(args);

// Add the "Controllers" service, which manages HTTP endpoints and API routes.
builder.Services.AddControllers();

// Add a service for API exploration and discovery.
builder.Services.AddEndpointsApiExplorer();

// Add a service for generating Swagger documentation for the API.
builder.Services.AddSwaggerGen();

// Build the web application based on the configured services.
var app = builder.Build();

// Check if the application is running in a development environment.
if (app.Environment.IsDevelopment())
{
    // Enable Swagger for API documentation.
    app.UseSwagger();

    // Enable Swagger UI for interactive API exploration.
    app.UseSwaggerUI();
}

// Enable HTTP to HTTPS redirection for secure communication.
app.UseHttpsRedirection();

// Handle authorization and security-related tasks.
app.UseAuthorization();

// Map the API controllers, making them accessible via HTTP routes.
app.MapControllers();

// Start listening for incoming HTTP requests.
app.Run();
