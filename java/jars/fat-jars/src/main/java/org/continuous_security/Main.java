package org.continuous_security;

import java.net.InetSocketAddress;
import java.util.List;
import java.util.ArrayList;

import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.socket.DatagramPacket;
import io.netty.handler.codec.dns.DnsResponseDecoder;
import io.netty.buffer.PooledByteBufAllocator;
import io.netty.buffer.ByteBufAllocator;
import io.netty.buffer.EmptyByteBuf;
import io.netty.buffer.ByteBuf;

public class Main {

  public static class MyDecoder extends DnsResponseDecoder {
    @Override
    public void decode(ChannelHandlerContext ctx, DatagramPacket packet, List<Object> out) throws Exception {
      super.decode(ctx, packet, out);
    }
  }

  public static void main(String[] args) {
    System.out.println("Started execution.");
    InetSocketAddress myAddress = new InetSocketAddress("localhost", 17385);

    ByteBufAllocator allocator = new PooledByteBufAllocator(true);
    ByteBuf buf = new EmptyByteBuf(allocator);

    DatagramPacket myPacket = new DatagramPacket(buf, myAddress);

    List<Object> myResponses = new ArrayList<Object>();

    MyDecoder decoder = new MyDecoder();
    try {
      // first argument is ChannelHandlerContext, but we will just pass in null to save trouble of instantiating it.
      decoder.decode(null, myPacket, myResponses);
    } catch (Exception e) {
      // ignored. There's probably an error with passing in null to decode()
      // but it doesn't matter.
    }
    System.out.println("Completed execution.");
  }
}
