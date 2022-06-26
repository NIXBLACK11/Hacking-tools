import java.net.*;
import java.net.http.HttpRequest;
import java.net.http.*;
import java.io.*;
public class mass 
{
    public static void main(String args[])
    {
        try
        {
            BufferedReader reader=new BufferedReader(new FileReader("masscan.txt"));
            String s;
            while((s=reader.readLine()) !=null)
            {      
                String h=check(s);
            }
            reader.close();
        }
        catch(IOException e)
        {
            e.printStackTrace();
        }
    }
    public static HttpResponse check(String domain)
    {
        try
        {
            domain = "https://"+domain;
            HttpClient client=HttpClient.newHttpClient();
            HttpRequest request=HttpRequest.newBuilder().uri(URI.create(domain)).build();
            HttpResponse response=client.send(request, HttpResponse.BodyHandlers.ofString());
            System.out.println(domain);
            return response;
        }
        catch(Exception e)
        {
            try
            {
                domain=domain.replace("https://","http://");
                HttpClient client=HttpClient.newHttpClient();
                HttpRequest request=HttpRequest.newBuilder().uri(URI.create(domain)).build();
                HttpResponse response=client.send(request, HttpResponse.BodyHandlers.ofString());
                System.out.println(domain);
                return response;
            }
            catch(Exception e2)
            {
                return null;
            }
        }
    }
}
