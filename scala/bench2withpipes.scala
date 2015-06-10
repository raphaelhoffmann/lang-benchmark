import java.io._
import scala.collection._

object bench2withpipes {

  val set = new mutable.HashSet[String]()

  def main(args:Array[String]) {
    readDict

    val reader = new BufferedReader(new InputStreamReader(System.in), 1024*1024)
    val writer = new BufferedWriter(new OutputStreamWriter(System.out), 1024*1024)

    var l:String = null
    while ({ l = reader.readLine; l != null }) {
      val toks = l.split(" ")
      for (tok <- toks)
        if (set.contains(tok))
          writer.write("match\n")
    }
    reader.close
    writer.close
  }

  def readDict = {
    val reader = new BufferedReader(new InputStreamReader(new FileInputStream("../data/english_words.tsv")))
    var l:String = null
    while ({ l = reader.readLine; l != null })
      set += l
    reader.close
  }

}
