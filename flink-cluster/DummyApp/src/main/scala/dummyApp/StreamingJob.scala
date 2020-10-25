package dummyApp

import org.apache.flink.streaming.api.scala._

object StreamingJob {
  def main(args: Array[String]) {
    val env = StreamExecutionEnvironment.getExecutionEnvironment
    env.fromCollection(List(0).map(x => {
      println(s"Sleeping for ${Long.MaxValue}ms")
      Thread.sleep(Long.MaxValue)
    }))
    env.execute("Example app - infinite no-op stream")
  }
}
