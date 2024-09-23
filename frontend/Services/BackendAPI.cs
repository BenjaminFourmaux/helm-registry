using frontend.Models;
using System.Text.Json;

namespace frontend.Services
{
    public class BackendAPI
    {
        private readonly HttpClient _client;
        public RegistryInfo RegistryInfo { get; private set; }

        public BackendAPI(IHttpClientFactory factory) 
        {
            _client = factory.CreateClient("backendAPI");
            RegistryInfo = new RegistryInfo();
        }

        public async Task GetRegistryInfoAsync()
        {
            Console.WriteLine("Getting registry info");
            this.RegistryInfo = await GetHomeData();
        }

        public async Task<RegistryInfo> GetHomeData()
        {
            var response = await _client.GetAsync("/bff/home");
            response.EnsureSuccessStatusCode();

            var content = await response.Content.ReadAsStringAsync();
            Console.WriteLine(content);
            return JsonSerializer.Deserialize<RegistryInfo>(content, new JsonSerializerOptions{ PropertyNameCaseInsensitive = true,})!;;
        }

        public async void Ping()
        {
            await _client.GetAsync("/");
        }
    }
}
