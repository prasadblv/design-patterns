#!/usr/bin/env python3

from abc import ABC, abstractmethod

class Notification(ABC):

    @abstractmethod
    def send(self, msg : str) -> None:
        pass

class EmailNotification(Notification):

    def send(self, msg: str) -> None:
        print("Sending Email..." + msg)

class SMSNotifictaion(Notification):

    def send(self, msg : str) -> None:
        print("Sending sms..." + msg)


class NotificationFactory:

    @staticmethod
    def create(type: str) -> Notification:
        match type:
            case "email":
                return EmailNotification() 
            case "sms":
                return SMSNotifictaion()

if __name__ == "__main__":
    email_notification = NotificationFactory.create("email").send("email")
    sms_notification = NotificationFactory.create("sms").send("sms")


