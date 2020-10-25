/*
 * Example copied from here:
 * https://github.com/apache/flink/blob/master/flink-examples/flink-examples-streaming/src/main/scala/org/apache/flink/streaming/scala/examples/async/AsyncIOExample.scala
 */
package dummyApp

import java.util.concurrent.TimeUnit

import org.apache.flink.streaming.api.functions.source.ParallelSourceFunction
import org.apache.flink.streaming.api.functions.source.SourceFunction.SourceContext
import org.apache.flink.streaming.api.scala._
import org.apache.flink.streaming.api.scala.async.ResultFuture

import scala.concurrent.{ExecutionContext, Future}

object StreamingJob {

  def main(args: Array[String]) {
    val timeout = 10000L

    val env = StreamExecutionEnvironment.getExecutionEnvironment

    val input = env.addSource(new SimpleSource())

    val asyncMapped = AsyncDataStream.orderedWait(input, timeout, TimeUnit.MILLISECONDS, 10) {
      (input, collector: ResultFuture[Int]) =>
        Future {
          collector.complete(Seq(input))
        } (ExecutionContext.global)
    }

    asyncMapped.print()

    env.execute("Example app - Infinite counter")
  }
}

class SimpleSource extends ParallelSourceFunction[Int] {
  var running = true
  var counter = 0

  override def run(ctx: SourceContext[Int]): Unit = {
    while (running) {
      ctx.getCheckpointLock.synchronized {
        ctx.collect(counter)
      }
      counter += 1

      Thread.sleep(1000L)
    }
  }

  override def cancel(): Unit = {
    running = false
  }
}