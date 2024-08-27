using frontend.Models;
using System.Text.Json;

namespace frontend.Services
{
    public class BackendAPI
    {
        private readonly HttpClient _client;

        public BackendAPI(IHttpClientFactory factory) 
        {
            _client = factory.CreateClient("backendAPI");
        }

        public async Task<BackendAPIHomeResponse> GetHomeData()
        {
            var response = await _client.GetAsync("/bff/home");
            response.EnsureSuccessStatusCode();

            var content = await response.Content.ReadAsStringAsync();
            Console.WriteLine(content);
            return JsonSerializer.Deserialize<BackendAPIHomeResponse>(content, new JsonSerializerOptions{ PropertyNameCaseInsensitive = true,})!;;
        }

        public async void Ping()
        {
            await _client.GetAsync("/");
        }
    }
}
