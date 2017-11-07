# Python Project with requirements.txt

This project is an example of a python project that uses requirements.txt to manage its dependencies.

`requirements.txt` at the root directory declares the dependencies:
- django==1.7.1
- PyJWT==0.4.2
- rsa==3.0
- requests==2.2.1
- feedparser==5.1.1
- pycrypto==2.4
- Flask==0.11.1

Total number of dependencies:
- 7 direct dependencies.
- 6 transitive dependencies.
- 33 vulnerabilities.

#### Vulnerable Call Chain
`main.py` has a method call that leads to a vulnerability.

#### What's next?
By running the security scanner of your choice against `with-requirements-file` project, you should get results that are coherent with the details above if the tool supports using requirements.txt to manage python dependencies.
