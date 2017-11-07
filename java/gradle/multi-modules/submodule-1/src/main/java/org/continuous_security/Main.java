package org.continuous_security;

import org.mindrot.jbcrypt.BCrypt;

public class Main {

  public static void main(String[] args) {
    String candidate = args[0];
    String hashed = BCrypt.hashpw(candidate, BCrypt.gensalt(12));
    BCrypt.checkpw(candidate, hashed);
    System.out.println("Gradle multi-module submodule-1 completed. Args: " + args[0]);
  }
}
