using frontend.Client;
using frontend.Services;
using System.Reflection;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.Services.AddRazorComponents()
    .AddInteractiveServerComponents();

// Declare HTTP clients
builder.Services.AddHttpClient("backendAPI", config =>
{
    config.BaseAddress = new Uri("http://localhost:8080/");
});

builder.Services.AddSingleton<BackendAPI>();

// Add BlazorBootstrap service
builder.Services.AddBlazorBootstrap();
Console.WriteLine("Using BlazorBootstrap version: " + Assembly.GetAssembly(typeof(BlazorBootstrap.TypeExtensions)).GetName().Version.ToString());

var app = builder.Build();

// Configure the HTTP request pipeline.
if (!app.Environment.IsDevelopment())
{
    app.UseExceptionHandler("/Error", createScopeForErrors: true);
    // The default HSTS value is 30 days. You may want to change this for production scenarios, see https://aka.ms/aspnetcore-hsts.
    app.UseHsts();
}

app.UseHttpsRedirection();

app.UseStaticFiles();
app.UseAntiforgery();

app.MapRazorComponents<App>()
    .AddInteractiveServerRenderMode();

app.Run();
