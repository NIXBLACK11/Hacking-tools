import java.net.*;
import java.net.http.HttpRequest;
import java.net.http.*;
import java.io.*;
public class toex extends Thread
{
    private String domain;
    public toex(String domain)
    {
        this.domain=domain;
    }
    public void run()
    {
        try
        {
            HttpResponse response=check_alive();
            domain=domain.replace("://","_");
            BufferedWriter writer1=new BufferedWriter(new FileWriter("out//index.txt",true));     
            writer1.write(response.request().toString().split(" ")[0]+" "+response.body().toString().split("title")[1].substring(1,response.body().toString().split("title")[1].length()-2)+" ("+response.statusCode()+") "+" ["+response.body().toString().length()+"]"+"\n");
            writer1.close();
            BufferedWriter writer2=new BufferedWriter(new FileWriter("out//header//"+domain+".txt"));     
            writer2.write(response.headers()+"\n");
            writer2.close();
            BufferedWriter writer3=new BufferedWriter(new FileWriter("out//body//"+domain+".txt"));     
            writer3.write(response.body()+"\n");
            writer3.close();
            /*BufferedWriter writer4=new BufferedWriter(new FileWriter("out//java_script//"+domain+".txt"));     
            String s=response.body().toString();
            String js[]=java_script(s);
            for(int o=0;o<c;o++)
            {
                writer4.write("\n"+js[0]);
            }
            writer4.close();*/
        }
        catch(Exception e){}
    }
    public HttpResponse check_alive()
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