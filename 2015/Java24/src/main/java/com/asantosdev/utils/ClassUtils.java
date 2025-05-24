package com.asantosdev.utils;

import java.io.IOException;
import java.net.URISyntaxException;
import java.net.URL;
import java.nio.file.Files;
import java.nio.file.Paths;

public class ClassUtils {

  public static String getResourceAsString(String resourcePath) {
    try {

      URL inputResourceURL = ClassUtils.class.getClassLoader().getResource(resourcePath);
      assert inputResourceURL != null;
      return Files.readString(Paths.get(inputResourceURL.toURI()));

    } catch (URISyntaxException e) {
      String message = String.format("Path URI malformed: %s", resourcePath);
      throw new RuntimeException(message, e);
    } catch (IOException e) {
      String message = String.format("Failed to read resource: %s", resourcePath);
      throw new RuntimeException(message, e);
    }
  }
}
