package org.continuous_security;

import org.apache.commons.fileupload.MultipartStream;
import java.io.ByteArrayInputStream;

public class Main {

  private static void filterXMLSignature() throws Exception {
    byte[] bytes = new byte[256];
    new MultipartStream(new ByteArrayInputStream(bytes), bytes);
  }

  public static void main(String[] args) throws Exception {
    filterXMLSignature();
  }
}
