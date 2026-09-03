package creational.factory;

interface Notification{
    void send(String message);
}

class EmailNotification implements Notification{

    @Override
    public void send(String msg){
        System.out.println(msg);
    }
}

class SmsNotification implements Notification{
    
    @Override
    public void send(String sms){
        System.out.println(sms);
    }
}

class DefaultNotification implements Notification{
    @Override
    public void send(String msg){
        System.out.println(msg);
    }
}

class NotificationFactory{
    public static Notification create(String type){
        Notification notif = null;
        switch(type){
            case "email":
                notif = new EmailNotification();
            break;
            
            case "sms":
                notif = new SmsNotification();
            break;
            
            default:
                notif = new DefaultNotification();
            break;    
        }
        return notif;
    }
}

public class Main{

    public static void main(String[] args) {
        NotificationFactory.create("email").send("email");
        NotificationFactory.create("sms").send("sms");
        NotificationFactory.create("def").send("default");
    }
}