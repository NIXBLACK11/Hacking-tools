import java.io.*;
public class htoo 
{
    public static void main(String args[])throws IOException
    {
        BufferedReader buff=new BufferedReader(new InputStreamReader(System.in));
        System.out.println("enter the file");
        String f=buff.readLine();
        new File("out").mkdir();
        new File("out//body").mkdir();
        new File("out//header").mkdir();
        new File("out//java_script").mkdir();
        try
        {
            BufferedReader reader=new BufferedReader(new FileReader(f));
            String dom;
            while((dom=reader.readLine()) !=null)
            {
                toex myThing=new toex(dom);
                myThing.start();
            }
             reader.close();
        }
        catch(IOException e)
        {
            System.out.println("NO such file");
        }
    }
}