# notifier

## What it is

**notifier** is a general purpose notification service to inform you via SMS, email, or telegram. It can be used by other services directly via REST API or indirectly via provided client. Currently support for only a few russian SMS services is available out of the box, but adding your own channel is trivial. Or you can drop me a line as well, and we'll see if support for your favorite SMS service can be added.  

## Features

- web SMS (beeline.ru, smsc.ru, websms.ru)
- telegram bot
- REST API
- command-line client
- easily extensible

## Example setup

### What we need

**notifyd** is written in Go, so you need to have some basic knowledge how to build Go programs. It shold build cleanly accorging to standard Go procedures once you install all the needed Go dependencies. There are not many of them.

**notifyd** is a web-application, hence it can be run on its own in *http* mode. But I'd recommend using it in *https* mode. For that we need a proxy (e.g **nginx**). Maybe sometimes I'll add *https* and other usefull stuff to **notifyd**, but right now we rely on proxy. A local mail server (e.g. **postfix**) is used to send out emails. We also need to get in touch with the **botfather** to register telegram bot.

I assume you know what an ssl sertificate is and how to handle it.

### nginx setup

We'll try to keep things as simple as possible. So here's the **nginx** configuration file *nginx.conf*:
```
events {
    worker_connections  1024;
}

error_log syslog:server=unix:/var/run/log;

http {

    map $ssl_client_s_dn $ssl_client_s_dn_cn {
        default "should_not_happen";
        ~/CN=(?<CN>[^/]+) $CN;
    }

    server {
        listen <--YOUR_IP_HERE-->:80;
        server_name <--SERVER_NAME_HERE-->;
        return 301 https://<--SERVER_NAME_HERE-->$request_uri;
    }

    server {

        access_log syslog:server=unix:/var/run/log;
        listen <--YOUR_IP_HERE-->:443 ssl;
        ssl_certificate     /etc/nginx/<--CERTIFICATE_FILE-->;
        ssl_certificate_key /etc/nginx/<--KEY_FILE-->;
        ssl_protocols       TLSv1 TLSv1.1 TLSv1.2;
        ssl_ciphers         HIGH:!aNULL:!MD5;

        server_name         <--SERVER_NAME_HERE-->;

        location / {
            proxy_pass http://localhost:8084/;
            proxy_set_header   X-Real-IP $remote_addr;
            proxy_set_header   Host $http_host;
            proxy_set_header   X-Forwarded-For $proxy_add_x_forwarded_for;
        }
    }
}
```
**nginx** speaks https to the outside world an communicates with **notifyd** via http on a special port.
It also sets the **X-Forwarded-For** header, just so **notifyd** knows the real address of the calling client.


### postfix setup

For **postfix** this minimal *main.cf* will do:
```
#
# LOCAL PATHNAME INFORMATION
#

sendmail_path = /usr/sbin/sendmail
newaliases_path = /usr/bin/newaliases
mailq_path = /usr/bin/mailq

html_directory = no
manpage_directory = /usr/share/man
readme_directory = no
queue_directory = /var/spool/postfix
command_directory = /usr/sbin
daemon_directory = /usr/libexec/postfix
data_directory = /var/run/postfix

#
# QUEUE AND PROCESS OWNERSHIP
#

mail_owner = postfix
setgid_group = postdrop

#
# NETWORK DETAILS
#

inet_protocols = ipv4
inet_interfaces = localhost

#
# LOCAL DELIVERY
#

# real users

home_mailbox = Maildir/
unknown_local_recipient_reject_code = 550
alias_maps = hash:/etc/postfix/aliases

# virtual users

virtual_mailbox_base=/
virtual_uid_maps=static:1002
virtual_gid_maps=static:1002

#
# TRUST AND RELAY CONTROL
#

#mynetworks = 168.100.189.0/28, 127.0.0.0/8
smtpd_recipient_restrictions = permit_mynetworks, permit_sasl_authenticated, reject_unauth_destination
strict_rfc821_envelopes = yes
```

This will launch **postfix** on 127.0.0.1, just what we need, because we only want so send out.

### telegram bot setup

### notifyd setup

## Invocation

### notifyd

### notify

## API

